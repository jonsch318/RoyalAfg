import React from "react";
import { SlotGame } from "./models/slot";

type SlotGameNullable = SlotGame | null;

const SlotContext = React.createContext<SlotGameNullable>(null);

export const SlotProvider = SlotContext.Provider;
export const SlotConsumer = SlotContext.Consumer;
export default SlotContext;
