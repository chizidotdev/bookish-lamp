export function formatDate(value: string) {
    const date = new Date(value);
    return date.toLocaleDateString('en-US', {
        year: '2-digit',
        month: 'short',
        day: 'numeric',
    });
}

export function getDateInput(value: string | Date) {
    const date = new Date(value);
    const day = date.getDate()
    const month = date.getMonth()
    const year = date.getFullYear()

    return date.toISOString().substring(0, 10);
}