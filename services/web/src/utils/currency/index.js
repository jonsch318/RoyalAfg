import Dinero from "dinero.js";

export function ToFormat(amount, currency = "USD") {
    return Dinero({ amount, currency }).toFormat();
}
