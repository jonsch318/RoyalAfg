export const SendEvent = (event) => {
    return JSON.stringify(event);
};

export const SendCreateEvent = (name, data) => {
    return SendEvent({ event: name, data: data });
};
