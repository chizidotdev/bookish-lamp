export function interceptor() {
    const { fetch: originalFetch } = window;

    window.fetch = async (...args) => {
        let [resource, config] = args;

        config = {
            ...config,
            headers: {
                ...config?.headers,
                'Content-Type': 'application/json',
            },
            credentials: 'include',
        };

        const response = await originalFetch(resource, config);
        return response;
    };
}
