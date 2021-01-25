export const usePokerTicketRequest = async (args) => {
    const params = new URLSearchParams({ buyIn: args.buyIn, class: args.class });
    let res = {};
    if (args.id) {
        res = await _fetch(`${process.env.NEXT_PUBLIC_POKER_TICKET_HOST}/api/poker/ticket/${args.id}`, params);
    } else {
        res = await _fetch(`${process.env.NEXT_PUBLIC_POKER_TICKET_HOST}/api/poker/ticket`, params);
    }
    return await res.json();
};

const _fetch = async (url, params) => {
    return fetch(`${url}?${params.toString()}`, {
        mode: "cors",
        credentials: "include",
        method: "GET"
    });
};
