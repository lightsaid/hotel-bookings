export function getImageURL(url: string | undefined) {
    if (!url) {
        return url;
    }
    if (url.indexOf("http://") === 0 || url.indexOf("https://") === 0) {
        return url;
    }

    return `${import.meta.env.VITE_BASE_API}${url}`;
}
