import { Issue } from "../lib/interfaces/issue.interface";

const index = async (): Promise<Issue[]> => {
    const response = await fetch("/api/issue");
    return await response.json() as Issue[];
}

export default {
    index
}