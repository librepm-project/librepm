import { en as baseEn } from 'vuetify/locale';

export const en = {
    global: {
        select: 'Select',
        search: 'Search',
        create: 'Create',
        name: 'Name',
        description: 'Description',
        public: 'Public',
        edit: 'Edit',
        delete: 'Delete',
        update: 'Update',
        delete_confirm: 'Are you sure you want to delete this item?',
        actions: 'Actions',
        details: 'Details',
        no_data: 'No data',
    },
    title: {
        issueIndex: 'Issues',
        issueCreate: "Create issue",
        adminProjectIndex: 'Projects',
        adminBoardIndex: 'Boards',
        dashboardShow: 'Dashboards',
        boardShow: 'Board',
        filterIndex: 'Filters',
        adminStatusIndex: 'Statuses',
        adminTrackerIndex: 'Trackers',
        issueShow: 'Issue details',
        login: 'Login',
        projectCreate: 'Create project',
        dashboardCreate: 'Create dashboard',
        boardCreate: 'Create board',
        filterCreate: 'Create filter',
        adminTrackerCreate: 'Create tracker',
        adminStatusCreate: 'Create status',
        adminUserIndex: 'Users',
        adminStatusShow: 'Edit Status',
        adminTrackerShow: 'Edit Tracker',
        adminProjectShow: 'Edit Project',
        adminUserShow: 'Edit User',
    },
    issue: {
        description: 'Description',
        summary: 'Summary',
        status: 'Status',
        key: 'Key',
        project: 'Project',
        tracker: 'Tracker',
    },
    board: {
        boards: 'Boards',
        columns: 'Columns',
        add_column: 'Add Column',
        column_label: 'Column Label',
        column_statuses: 'Statuses',
        no_columns_defined: 'No columns defined for this board.',
    },
    filter: {
        name: 'Name',
        description: 'Description',
        conditions: 'Filter Conditions',
        add_condition: 'Add Condition',
        field: 'Field',
        operator: 'Operator',
        value: 'Value',
    },
    dashboard: {
        add_widget: 'Add Widget',
        widget_type: 'Widget Type',
        widget_title: 'Title',
        select_filter: 'Select Filter',
        no_widgets: 'No widgets yet. Click "Add Widget" to get started.',
    },
    project: {
        codeName: 'Code name',
    },
    status: {
        name: 'Name',
        color: 'Color',
    },
    tracker: {
        name: 'Name',
        color: 'Color',
    },
    user: {
        email: 'E-mail',
        password: 'Password',
        fullName: 'Full name',
        firstName: 'First name',
        lastName: 'Last name',
        phone: 'Phone',
        language: 'Language',
        country: 'Country',
    },
    login: {
        submit: 'Login'
    },
    rule: {
        email: "Invalid email format",
        required: "Required field",
        minimumLength: "Minimum length is {len}"
    }
}

// Merge with Vuetify's English locale
Object.assign(en, baseEn);