import { ComponentProps } from 'react';

export interface ButtonProps
    extends Omit<ComponentProps<'button'>, 'className'> {
    children: React.ReactNode;
}

export function Button({ children }: ButtonProps) {
    return <button className="btn">{children}</button>;
}

export default Button;
