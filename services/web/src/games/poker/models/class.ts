export interface IClass {
    min: number;
    max: number;
    blind: number;
}

export interface ILobby {
    id: string;
    class: IClass;
    classIndex: number;
    i: number;
    changeClass: boolean;
}

export const LobbyInit: ILobby = {
    i: -1,
    classIndex: -1,
    class: {
        min: 0,
        max: 0,
        blind: 0
    },
    id: "",
    changeClass: false
};
