import { Fragment } from 'react';
import { Combobox, Transition } from '@headlessui/react';
import { ArrowUpDownIcon, CheckIcon } from '@chakra-ui/icons';
import { Box, FormLabel } from '@chakra-ui/react';
import { Input } from '~components';
import { Item } from '~lib/types';
import { UseSelectSalesItemReturnType } from '~modules/sales/select-sales-item/useSelectSalesItem';

type SelectSalesItemProps = UseSelectSalesItemReturnType;

export function SelectSalesItem(props: SelectSalesItemProps) {
    const { setSelected, selected, filteredItems, setQuery, query } =
        props;

    return (
        <Box>
            <Combobox value={selected} onChange={setSelected}>
                <FormLabel>Select Item</FormLabel>
                <div className='relative mt-1'>
                    <div>
                        <Combobox.Input
                            as={Input}
                            displayValue={(item: Item) => item.title}
                            onChange={(event) =>
                                setQuery(event.target.value)
                            }
                        />
                        <Combobox.Button className='absolute inset-y-0 right-0 flex items-center pr-2'>
                            <ArrowUpDownIcon
                                className='h-5 w-5 text-gray-400'
                                aria-hidden='true'
                            />
                        </Combobox.Button>
                    </div>
                    <Transition
                        as={Fragment}
                        leave='transition ease-in duration-100'
                        leaveFrom='opacity-100'
                        leaveTo='opacity-0'
                        afterLeave={() => setQuery('')}
                    >
                        <Combobox.Options className='absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm'>
                            {filteredItems.length === 0 &&
                            query !== '' ? (
                                <div className='relative cursor-default select-none py-2 px-4 text-gray-700'>
                                    Nothing found.
                                </div>
                            ) : (
                                filteredItems.map((item) => (
                                    <Combobox.Option
                                        key={item.id}
                                        className={({ active }) =>
                                            `relative cursor-default select-none py-2 pl-10 pr-4 ${
                                                active
                                                    ? 'bg-sky-700 text-white'
                                                    : 'text-gray-900'
                                            }`
                                        }
                                        value={item}
                                    >
                                        {({ selected, active }) => (
                                            <>
                                                <span
                                                    className={`block truncate ${
                                                        selected
                                                            ? 'font-medium'
                                                            : 'font-normal'
                                                    }`}
                                                >
                                                    {item.title}
                                                </span>
                                                {selected ? (
                                                    <span
                                                        className={`absolute inset-y-0 left-0 flex items-center pl-3 ${
                                                            active
                                                                ? 'text-white'
                                                                : 'text-teal-600'
                                                        }`}
                                                    >
                                                        <CheckIcon
                                                            className='h-5 w-5'
                                                            aria-hidden='true'
                                                        />
                                                    </span>
                                                ) : null}
                                            </>
                                        )}
                                    </Combobox.Option>
                                ))
                            )}
                        </Combobox.Options>
                    </Transition>
                </div>
            </Combobox>
        </Box>
    );
}

export default SelectSalesItem;
