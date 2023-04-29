import { Button } from '@copia/ui';
import { Dialog, Transition } from '@headlessui/react';
import { Fragment } from 'react';
import { useMutation, useQueryClient } from 'react-query';
import { useNavigate, useParams } from 'react-router-dom';
import { deleteItem } from './api';

export function DeleteItem() {
    const { id } = useParams();
    const navigate = useNavigate();
    const queryClient = useQueryClient();
    const { mutate, isLoading } = useMutation({
        mutationFn: deleteItem,
        onSuccess: () => {
            navigate('/');
            queryClient.invalidateQueries(['get-items']);
        },
    });

    function handleDelete() {
        if (!id) return;
        mutate(id);
    }

    return (
        <Transition appear show as={Fragment}>
            <Dialog
                as="div"
                className="relative z-10"
                onClose={() => navigate('/')}
            >
                <Transition.Child
                    as={Fragment}
                    enter="ease-out duration-300"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in duration-200"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                >
                    <div className="fixed inset-0 bg-black bg-opacity-40" />
                </Transition.Child>

                <div className="fixed inset-0 overflow-y-auto">
                    <div className="flex min-h-full items-center justify-center p-4">
                        <Transition.Child
                            as={Fragment}
                            enter="ease-out duration-300"
                            enterFrom="opacity-0 scale-95"
                            enterTo="opacity-100 scale-100"
                            leave="ease-in duration-200"
                            leaveFrom="opacity-100 scale-100"
                            leaveTo="opacity-0 scale-95"
                        >
                            <Dialog.Panel className="modal-box">
                                <Dialog.Title
                                    as="h3"
                                    className="text-lg font-medium leading-6"
                                >
                                    Confirm
                                </Dialog.Title>
                                <div className="mt-2">
                                    <p className="text-sm text-gray-500">
                                        Are you sure you want to delete this
                                        item. This cannot be undone.
                                    </p>
                                </div>

                                <div className="flex justify-end mt-4">
                                    <Button
                                        size="small"
                                        variant="danger"
                                        type="button"
                                        loading={isLoading}
                                        onClick={handleDelete}
                                    >
                                        Delete
                                    </Button>
                                </div>
                            </Dialog.Panel>
                        </Transition.Child>
                    </div>
                </div>
            </Dialog>
        </Transition>
    );
}
