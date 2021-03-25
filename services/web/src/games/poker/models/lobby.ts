export interface ILobby {
    lobbyId: string;
    blind: number;
    minBuyIn: number;
    maxBuyIn: number;
    minPlayersToStart: number;
    playerCount: number;
    gameStartTimeout: number;
}
