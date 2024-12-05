import { Project } from "@/lib/interfaces/project.interface";
import { Status } from "@/lib/interfaces/status.interface";
import { Tracker } from "@/lib/interfaces/tracker.interface";

export interface Issue {
    id: string;
    key: string;
    tracker: Tracker;
    summary: string;
    description: string;
    status: Status; 
    project: Project;
}