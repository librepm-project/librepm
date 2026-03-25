import { RelatedIssue } from "@/lib/interfaces/related-issue.interface";
import api from "@/api/api";

const getRelated = async (issueId: string, type?: string): Promise<RelatedIssue[]> => {
    const response = await api.apiCall().get(`/issue/${issueId}/related`, { params: { type } });
    return response.data as RelatedIssue[];
}

const create = async (issueId: string, relation: { targetIssueId: string; type: string }): Promise<RelatedIssue> => {
    const response = await api.apiCall().post(`/issue/${issueId}/related`, relation);
    return response.data as RelatedIssue;
}

const destroy = async (relatedIssueId: string): Promise<void> => {
    await api.apiCall().delete(`/related-issue/${relatedIssueId}`);
}

export default {
    getRelated,
    create,
    destroy
}
