import { useState } from 'react';
import { Combobox } from '@headlessui/react';
import { BsCheck2, BsChevronExpand } from 'react-icons/bs';
import { Input } from '../input/input';

type ComboBoxData = { name: string; ID: string };
type ComboBoxProps<T extends ComboBoxData> = {
    data: T[] | undefined;
    label?: string;
    placeholder?: string;
    selected: T | undefined;
    setSelected: React.Dispatch<T>;
};

export function ComboBox<T extends ComboBoxData>({
    data,
    selected,
    setSelected,
}: ComboBoxProps<T>) {
    const [query, setQuery] = useState('');

    const filteredItems =
        query === ''
            ? data
            : data?.filter((item) =>
                  item.name
                      .toLowerCase()
                      .includes(query.toLowerCase().replace(/\s+/g, ''))
              );

    return (
        <Combobox
            as="div"
            className="relative"
            value={selected}
            onChange={setSelected}
        >
            <div className="relative w-full">
                <Combobox.Input
                    as={Input}
                    onChange={(event) => setQuery(event.target.value)}
                    placeholder="Search items..."
                    value={selected?.name}
                />
                <Combobox.Button className="absolute inset-y-0 right-0 flex items-center pr-4">
                    <BsChevronExpand
                        className="h-5 w-5 text-gray-400"
                        aria-hidden="true"
                    />
                </Combobox.Button>
            </div>

            <Combobox.Options className="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-base-200 py-1 text-base shadow-lg sm:text-sm">
                {filteredItems?.length === 0 && query !== '' && (
                    <div className="relative cursor-default select-none py-2 px-4">
                        Nothing found.
                    </div>
                )}

                {filteredItems?.map((item) => (
                    <Combobox.Option
                        key={item.ID}
                        value={item}
                        className={({ active }) =>
                            `relative cursor-default select-none py-2 pl-10 pr-4 ${
                                active ? 'bg-primary' : ''
                            }`
                        }
                    >
                        {({ selected }) => (
                            <>
                                <span
                                    className={`block truncate ${
                                        selected ? 'font-bold' : 'font-normal'
                                    }`}
                                >
                                    {item.name}
                                </span>
                                {selected ? (
                                    <span
                                        className={`absolute inset-y-0 left-0 flex items-center pl-3`}
                                    >
                                        <BsCheck2
                                            className="h-5 w-5"
                                            aria-hidden="true"
                                        />
                                    </span>
                                ) : null}
                            </>
                        )}
                    </Combobox.Option>
                ))}
            </Combobox.Options>
        </Combobox>
    );
}
