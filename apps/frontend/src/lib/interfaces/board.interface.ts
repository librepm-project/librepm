import { Status } from "./status.interface";

export interface BoardColumn {
    id?: string;
    label: string;
    weight?: number;
    statuses: Status[]; // This should match the backend's StatusResponse structure
    // For Request
    statusIds?: string[];
}

export interface Board {
    id?: string;
    name: string;
    description: string;
    boardColumns: BoardColumn[];
};
