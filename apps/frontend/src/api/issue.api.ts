import { Issue } from "../lib/interfaces/issue.interface";

const indexIssue = async (): Promise<Issue[]> => {
    const response = await fetch("/api/issue");
    return await response.json() as Issue[];
}

export default {
    indexIssue
}