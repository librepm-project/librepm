export interface Release {
  id?: string;
  name: string;
  status: string;
  released_at?: string | null;
  created_at?: string;
  updated_at?: string;
}
