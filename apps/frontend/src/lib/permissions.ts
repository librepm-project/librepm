export const Permissions = {
  IssueCreate:  'issue:create',
  IssueRead:    'issue:read',
  IssueUpdate:  'issue:update',
  IssueDelete:  'issue:delete',

  ProjectCreate: 'project:create',
  ProjectRead:   'project:read',
  ProjectUpdate: 'project:update',
  ProjectDelete: 'project:delete',

  CommentCreate: 'comment:create',
  CommentRead:   'comment:read',
  CommentUpdate: 'comment:update',
  CommentDelete: 'comment:delete',

  WorklogCreate: 'worklog:create',
  WorklogRead:   'worklog:read',
  WorklogUpdate: 'worklog:update',
  WorklogDelete: 'worklog:delete',

  AttachmentCreate: 'attachment:create',
  AttachmentRead:   'attachment:read',
  AttachmentDelete: 'attachment:delete',

  BoardCreate: 'board:create',
  BoardRead:   'board:read',
  BoardUpdate: 'board:update',
  BoardDelete: 'board:delete',

  FilterCreate: 'filter:create',
  FilterRead:   'filter:read',
  FilterUpdate: 'filter:update',
  FilterDelete: 'filter:delete',

  DashboardCreate: 'dashboard:create',
  DashboardRead:   'dashboard:read',
  DashboardUpdate: 'dashboard:update',
  DashboardDelete: 'dashboard:delete',

  ReleaseCreate: 'release:create',
  ReleaseRead:   'release:read',
  ReleaseUpdate: 'release:update',
  ReleaseDelete: 'release:delete',

  NotificationRead:   'notification:read',
  NotificationUpdate: 'notification:update',

  StatusCreate:  'status:create',
  StatusRead:    'status:read',
  StatusUpdate:  'status:update',
  StatusDelete:  'status:delete',

  TrackerCreate: 'tracker:create',
  TrackerRead:   'tracker:read',
  TrackerUpdate: 'tracker:update',
  TrackerDelete: 'tracker:delete',

  PriorityCreate: 'priority:create',
  PriorityRead:   'priority:read',
  PriorityUpdate: 'priority:update',
  PriorityDelete: 'priority:delete',

  TransitionCreate: 'transition:create',
  TransitionRead:   'transition:read',
  TransitionUpdate: 'transition:update',
  TransitionDelete: 'transition:delete',

  UserCreate: 'user:create',
  UserRead:   'user:read',
  UserUpdate: 'user:update',
  UserDelete: 'user:delete',

  SettingRead:   'setting:read',
  SettingUpdate: 'setting:update',
} as const;

export type Permission = typeof Permissions[keyof typeof Permissions];
