import { cva, type VariantProps } from 'class-variance-authority';
import ErrorIcon from './error-icon';
import WarningIcon from './warning-icon';

export type AlertProps = VariantProps<typeof alertStyles> & {
    children?: React.ReactNode;
    message: string;
};

const alertStyles = cva('alert shadow-lg', {
    variants: {
        variant: {
            warning: ['alert-warning'],
            error: ['alert-error'],
        },
    },
    defaultVariants: {
        variant: 'warning',
    },
});

export function Alert({ children, message, ...props }: AlertProps) {
    return (
        <div className={alertStyles(props)}>
            <div>
                {icons[props.variant ?? 'warning']}
                <span>{message}</span>
            </div>
            <div className="flex-none">{children}</div>
        </div>
    );
}

export default Alert;

const icons = {
    error: <ErrorIcon />,
    warning: <WarningIcon />,
};
