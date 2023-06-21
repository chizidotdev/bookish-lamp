export function formatDate(value: string) {
    const date = new Date(value);
    return date.toLocaleDateString('en-US', {
        year: '2-digit',
        month: 'long',
        day: 'numeric',
    });
}
