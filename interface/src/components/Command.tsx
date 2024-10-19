/* eslint-disable max-len */

import React from 'react';
import { type DialogProps } from '@radix-ui/react-dialog';

import { Dialog, DialogContent } from '@monetr/interface/components/Dialog';
import mergeTailwind from '@monetr/interface/util/mergeTailwind';

import { Command as CommandPrimitive } from 'cmdk';
import { Search } from 'lucide-react';
 
const Command = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive>
>(({ className, ...props }, ref) => (
  <CommandPrimitive
    ref={ ref }
    className={ mergeTailwind(
      'flex h-full w-full flex-col overflow-hidden rounded-lg bg-dark-monetr-background text-dark-monetr-content-emphasis',
      className
    ) }
    { ...props }
  />
));
Command.displayName = CommandPrimitive.displayName;
 
interface CommandDialogProps extends DialogProps {}
 
const CommandDialog = ({ children, ...props }: CommandDialogProps) => {
  return (
    <Dialog { ...props }>
      <DialogContent className='overflow-hidden p-0 shadow-lg'>
        <Command className='[&_[cmdk-group-heading]]:px-2 [&_[cmdk-group-heading]]:font-semibold [&_[cmdk-group-heading]]:text-muted-foreground [&_[cmdk-group]:not([hidden])_~[cmdk-group]]:pt-0 [&_[cmdk-group]]:px-2 [&_[cmdk-input-wrapper]_svg]:h-5 [&_[cmdk-input-wrapper]_svg]:w-5 [&_[cmdk-input]]:h-12 [&_[cmdk-item]]:px-2 [&_[cmdk-item]]:py-3 [&_[cmdk-item]_svg]:h-5 [&_[cmdk-item]_svg]:w-5'>
          {children}
        </Command>
      </DialogContent>
    </Dialog>
  );
};
 
const CommandInput = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Input>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Input>
>(({ className, ...props }, ref) => (
  <div className='flex items-center border-b-[thin] border-dark-monetr-border px-3' cmdk-input-wrapper=''>
    <Search className='mr-2 h-4 w-4 shrink-0 opacity-50' />
    <CommandPrimitive.Input
      ref={ ref }
      className={ mergeTailwind(
        'flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-gray-400 disabled:cursor-not-allowed disabled:opacity-50',
        className
      ) }
      { ...props }
    />
  </div>
));
 
CommandInput.displayName = CommandPrimitive.Input.displayName;
 
const CommandList = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.List>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.List>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.List
    ref={ ref }
    className={ mergeTailwind('max-h-[300px] overflow-y-auto overflow-x-hidden', className) }
    { ...props }
  />
));
 
CommandList.displayName = CommandPrimitive.List.displayName;
 
const CommandEmpty = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Empty>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Empty>
>((props, ref) => (
  <CommandPrimitive.Empty
    ref={ ref }
    className='py-6 text-center text-sm'
    { ...props }
  />
));
 
CommandEmpty.displayName = CommandPrimitive.Empty.displayName;
 
const CommandGroup = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Group>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Group>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.Group
    ref={ ref }
    className={ mergeTailwind(
      'overflow-hidden p-1 text-dark-monetr-content-emphasis',
      '[&_[cmdk-group-heading]]:px-2 [&_[cmdk-group-heading]]:py-1.5',
      '[&_[cmdk-group-heading]]:text-xs',
      '[&_[cmdk-group-heading]]:font-semibold',
      '[&_[cmdk-group-heading]]:text-dark-monetr-content-muted',
      className
    ) }
    { ...props }
  />
));
 
CommandGroup.displayName = CommandPrimitive.Group.displayName;
 
const CommandSeparator = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Separator>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Separator>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.Separator
    ref={ ref }
    className={ mergeTailwind('-mx-1 h-px bg-border', className) }
    { ...props }
  />
));
CommandSeparator.displayName = CommandPrimitive.Separator.displayName;
 
const CommandItem = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Item>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.Item
    ref={ ref }
    className={ mergeTailwind([
      'text-dark-monetr-content-emphasis data-[disabled="true"]:text-dark-monetr-content-muted',
      'data-[disabled="false"]:cursor-pointer data-[disabled="true"]:cursor-default',
      'data-[disabled="true"]:pointer-events-none',
      'relative flex cursor-default select-none items-center rounded-sm',
      'px-2 py-1.5',
      'text-sm outline-none',
      'aria-selected:bg-dark-monetr-background-emphasis aria-selected:text-accent-foreground',
    ], className) }
    { ...props }
  />
));
 
CommandItem.displayName = CommandPrimitive.Item.displayName;
 
const CommandShortcut = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLSpanElement>) => {
  return (
    <span
      className={ mergeTailwind(
        'ml-auto text-xs tracking-widest text-muted-foreground',
        className
      ) }
      { ...props }
    />
  );
};
CommandShortcut.displayName = 'CommandShortcut';
 
export {
  Command,
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
  CommandShortcut,
};
