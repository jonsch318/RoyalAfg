export const getDateFromPast = (years: number): Date => {
    const date = new Date();
    date.setFullYear(date.getFullYear() - years);
    return date;
};
