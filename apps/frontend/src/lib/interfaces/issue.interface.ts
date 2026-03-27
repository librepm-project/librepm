import { Project } from "@/lib/interfaces/project.interface";
import { Status } from "@/lib/interfaces/status.interface";
import { Tracker } from "@/lib/interfaces/tracker.interface";
import { User } from "@/lib/interfaces/user.interface";

export interface Issue {
    id?: string;
    key?: string;
    summary: string;
    description: string;
    tracker?: Tracker;
    trackerId?: string;
    status?: Status;
    statusId?: string;
    project?: Project | null;
    projectId?: string;
    assignedUser?: User | null;
    assignedUserId?: string | null;
}