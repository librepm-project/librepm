import { Board } from "@/lib/interfaces/board.interface";
import api from "@/api/api";

const index = async (): Promise<Board[]> => {
    const response = await api.apiCall().get("/board");
    return response.data as Board[];
}

const show = async (boardId: string): Promise<Board> => {
    const response = await api.apiCall().get(`/board/${boardId}`);
    return response.data as Board;
}

const create = async (board: Partial<Board>): Promise<Board> => {
    const response = await api.apiCall().post(`/board`, board);
    return response.data as Board;
}

const update = async (boardId: string, board: Partial<Board>): Promise<Board> => {
    const response = await api.apiCall().put(`/board/${boardId}`, board);
    return response.data as Board;
}

const remove = async (boardId: string): Promise<void> => {
    await api.apiCall().delete(`/board/${boardId}`);
}

export default {
    index,
    show,
    create,
    update,
    remove
}
