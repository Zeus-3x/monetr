package recurring

import (
	"math"
	"sort"
	"time"

	"github.com/monetr/monetr/server/models"
)

type Detection struct {
	timezone           *time.Location
	preprocessor       *PreProcessor
	dbscan             *DBSCAN
	latestObservedDate time.Time
}

func NewRecurringTransactionDetection(timezone *time.Location) *Detection {
	return &Detection{
		timezone: timezone,
		preprocessor: &PreProcessor{
			documents: make([]Document, 0, 500),
			wc:        make(map[string]float64, 128),
		},
		dbscan:             nil,
		latestObservedDate: time.Time{},
	}
}

type RecurringTransaction struct {
	Name       string
	Window     WindowType
	Rule       *models.RuleSet
	First      time.Time
	Last       time.Time
	Next       time.Time
	Ended      bool
	Confidence float64
	Matches    []uint64
}

func (d *Detection) AddTransaction(txn *models.Transaction) {
	d.preprocessor.AddTransaction(txn)
	if txn.Date.After(d.latestObservedDate) {
		d.latestObservedDate = txn.Date
	}
}

func (d *Detection) GetRecurringTransactions() []RecurringTransaction {
	type Hit struct {
		Window WindowType
		Time   time.Time
	}
	type Miss struct {
		Window WindowType
		Time   time.Time
	}
	type Transaction struct {
		ID       uint64
		Name     string
		Merchant string
		Date     time.Time
		Amount   int64
	}

	d.dbscan = NewDBSCAN(d.preprocessor.GetDatums(), Epsilon, MinNeighbors)
	result := d.dbscan.Calculate()
	bestScores := make([]RecurringTransaction, 0, len(result))

	for _, cluster := range result {
		transactions := make([]*models.Transaction, 0, len(cluster.Items))
		for index := range cluster.Items {
			transactions = append(transactions, d.dbscan.dataset[index].Transaction)
		}
		sort.Slice(transactions, func(i, j int) bool {
			return transactions[i].Date.Before(transactions[j].Date)
		})

		start, end := transactions[0].Date, transactions[len(transactions)-1].Date
		windows := GetWindowsForDate(transactions[0].Date, d.timezone)
		scores := make([]RecurringTransaction, 0, len(windows))
		for _, window := range windows {
			misses := make([]Miss, 0)
			hits := make([]Hit, 0, len(transactions))
			ids := make([]uint64, 0, len(transactions))
			occurrences := window.Rule.Between(start.AddDate(0, 0, -window.Fuzzy), end.AddDate(0, 0, window.Fuzzy), false)
			for x := range occurrences {
				occurrence := occurrences[x]
				foundAny := false
				for i := range transactions {
					transaction := transactions[i]
					delta := math.Abs(transaction.Date.Sub(occurrence).Hours())
					fuzz := float64(window.Fuzzy) * 24
					if fuzz >= delta {
						foundAny = true
						hits = append(hits, Hit{
							Window: window.Type,
							Time:   occurrence,
						})
						ids = append(ids, transaction.TransactionId)
						continue
					}
				}
				if !foundAny {
					misses = append(misses, Miss{
						Window: window.Type,
						Time:   occurrence,
					})
				}
			}

			if len(hits) == 0 {
				continue
			}
			next := window.Rule.After(hits[len(hits)-1].Time, false)
			countHits := float64(len(hits))
			countMisses := float64(len(misses)) * 1.1
			countTxns := float64(len(transactions))
			ended := next.Before(d.latestObservedDate.AddDate(0, 0, -window.Fuzzy*2))
			latestTxn := transactions[len(transactions)-1]
			name := latestTxn.OriginalName
			if latestTxn.OriginalMerchantName != "" {
				name = latestTxn.OriginalMerchantName
			}

			scores = append(scores, RecurringTransaction{
				Name:       name,
				Window:     window.Type,
				Rule:       &models.RuleSet{Set: *window.Rule},
				First:      hits[0].Time,
				Last:       hits[len(hits)-1].Time,
				Next:       next,
				Ended:      ended,
				Confidence: (countHits - countMisses) / countTxns,
				Matches:    ids,
			})
		}

		sort.Slice(scores, func(i, j int) bool {
			return scores[i].Confidence > scores[j].Confidence
		})

		if scores[0].Confidence > 0.65 {
			bestScores = append(bestScores, scores[0])
		}
	}

	return bestScores
}
