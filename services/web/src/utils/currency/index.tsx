import Dinero from "dinero.js";

export function ToFormat(amount: number, currency: Dinero.Currency = "USD"): string {
    return Dinero({ amount, currency }).toFormat();
}
