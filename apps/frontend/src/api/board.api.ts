import { Board } from "../lib/interfaces/board.interface";

const indexBoard = async (): Promise<Board[]> => {
    const response = await fetch("/api/board");
    return await response.json() as Board[];
}

export default {
    indexBoard
}