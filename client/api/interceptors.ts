import axios from 'axios';

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

    axios.interceptors.request.use(
        (config) => {
            config.headers['Content-Type'] = 'application/json';
            config.withCredentials = true;
            return config;
        },
        (error) => {
            return Promise.reject(error);
        }
    );
}
