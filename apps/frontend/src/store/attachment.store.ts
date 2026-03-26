import { defineStore } from 'pinia';
import { Attachment } from '@/lib/interfaces/attachment.interface';
import attachmentApi from '@/api/attachment.api';

interface AttachmentStore {
  attachments: Attachment[];
}

export const useAttachmentStore = defineStore('attachment', {
  state: (): AttachmentStore => ({
    attachments: [],
  }),
  actions: {
    async getAttachments(issueId: string) {
      this.attachments = await attachmentApi.indexByIssue(issueId);
    },
    async create(issueId: string, file: File) {
      const created = await attachmentApi.createForIssue(issueId, file);
      this.attachments.unshift(created);
      return created;
    },
    async destroy(attachmentId: string) {
      await attachmentApi.destroy(attachmentId);
      this.attachments = this.attachments.filter(a => a.id !== attachmentId);
    },
  },
});
