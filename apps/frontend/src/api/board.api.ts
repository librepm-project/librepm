import { Board } from "../lib/interfaces/board.interface";
import api from "./api";

const index = async (): Promise<Board[]> => {
    const response = await api.apiCall().get("/api/board");
    return response.data as Board[];
}

const show = async (boardId: number): Promise<Board> => {
    const response = await api.apiCall().get(`/board/${boardId}`);
    return response.data as Board;
}

export default {
    index,
    show,
}