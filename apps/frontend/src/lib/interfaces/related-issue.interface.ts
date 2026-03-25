import { Issue } from "@/lib/interfaces/issue.interface";

export interface RelatedIssue {
    id?: string;
    issueA?: Issue;
    issueB?: Issue;
    type: 'related' | 'blocks' | 'implements' | 'duplicates';
    directionFromCurrent?: 'related' | 'blocks' | 'blocked_by' | 'implements' | 'implemented_by' | 'duplicates' | 'is_duplicated_by';
    createdAt?: string;
}
