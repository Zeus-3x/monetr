import React from 'react';
import { Meta, StoryObj } from '@storybook/react';
import { startOfTomorrow } from 'date-fns';

import MDatePicker from './MDatePicker';
import MSpan from './MSpan';
import MTextField from './MTextField';
import MForm from '@monetr/interface/components/MForm';


const meta: Meta<typeof MDatePicker> = {
  title: '@monetr/interface/components/Date Picker',
  component: MDatePicker,
};

export default meta;

export const Default: StoryObj<typeof MDatePicker> = {
  name: 'Default',
  render: () => (
    <MForm initialValues={ {} } onSubmit={ () => {} }>
      <div className='w-full flex p-4'>
        <div className='w-full max-w-xl grid grid-cols-1 grid-flow-row gap-1'>
          <MTextField 
            label="I'm a basic text field"
            labelDecorator={ () => <MSpan size='xs'>Just here for reference</MSpan> }
            placeholder='I can have some text here...' 
          />
          <MDatePicker
            label='When do you get paid next?'
            placeholder='Please select a date...'
            enableClear={ true }
          />
          <MDatePicker
            label='When do you get paid next?'
            placeholder='A really really really really really long placeholder that should be bigger than the field...'
            enableClear={ true }
          />
          <MDatePicker
            label='Must be in the future.'
            placeholder='Please select a date...'
            min={ startOfTomorrow() }
          />
          <MDatePicker
            label='Go by year'
            placeholder='Please select a date...'
            enableYearNavigation
          />
          <MDatePicker
            label='Required date picker'
            placeholder='You must select a date...'
            required
          />
          <MDatePicker
            label='Disabled date picker'
            placeholder='You cannot select a date...'
            enableYearNavigation
            disabled
          />
          <MDatePicker
            label='With an error!'
            placeholder='Please select a date...'
            error='Invalid date selected!'
          />
        </div>
      </div>
    </MForm>
  ),
};
