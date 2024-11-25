import { Project } from "../lib/interfaces/project.interface";

const index = async (): Promise<Project[]> => {
    const response = await fetch("/api/project");
    return await response.json() as Project[];
}

export default {
    index
}