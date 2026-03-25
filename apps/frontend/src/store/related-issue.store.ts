import { defineStore } from 'pinia';
import { RelatedIssue } from '@/lib/interfaces/related-issue.interface';
import relatedIssueApi from '@/api/related-issue.api';

interface RelatedIssueStore {
    relations: RelatedIssue[];
}

export const useRelatedIssueStore = defineStore('relatedIssue', {
    state: (): RelatedIssueStore => ({
        relations: [],
    }),
    actions: {
        async getRelated(issueId: string, type?: string) {
            this.relations = await relatedIssueApi.getRelated(issueId, type);
        },
        async create(issueId: string, targetIssueId: string, type: string) {
            const created = await relatedIssueApi.create(issueId, { targetIssueId, type });
            this.relations.push(created);
            return created;
        },
        async delete(relatedIssueId: string) {
            await relatedIssueApi.destroy(relatedIssueId);
            this.relations = this.relations.filter(r => r.id !== relatedIssueId);
        }
    }
});
