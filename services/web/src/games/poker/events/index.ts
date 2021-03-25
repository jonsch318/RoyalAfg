import { IEvent } from "../models/event";

export const SendEvent = (event: IEvent) => {
    return JSON.stringify(event);
};
