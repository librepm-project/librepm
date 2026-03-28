export interface Project {
    id?: string;
    name: string;
    codeName: string;
    defaultStatusId?: string | null;
    defaultTrackerId?: string | null;
}