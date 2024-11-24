import { Project } from "../lib/interfaces/project.interface";

const indexProject = async (): Promise<Project[]> => {
    const response = await fetch("/api/project");
    return await response.json() as Project[];
}

export default {
    indexProject
}