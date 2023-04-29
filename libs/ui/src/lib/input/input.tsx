import { ComponentProps, forwardRef } from 'react';

export type InputProps = Omit<ComponentProps<'input'>, 'className'> & {
    label?: string;
};
export type InputRef = HTMLInputElement;

export const Input = forwardRef<InputRef, InputProps>(
    ({ label, ...props }, ref) => {
        return (
            <div className="form-control w-full">
                {label && (
                    <label className="label">
                        <span className="label-text">{label}</span>
                    </label>
                )}
                <input
                    ref={ref}
                    className="input input-bordered w-full"
                    {...props}
                />
            </div>
        );
    }
);

export default Input;
