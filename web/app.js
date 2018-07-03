var user = {
    ID: 0,
    Login: '',
    Password: '',
    Username: '',
    Email: '',
    Status: false,
    UserIcon: ''
};

var message = {
    User: user,
    Group: { GroupName: '' },
    Content: ''
};
var group = {
    GroupName: '',
    Messages: [message],
    Members: [user],
    ID: null
};

var test = new Vue({
    el: '#app',
    data: {
        MessageIn: {
            User: user,
            Contact: user,
            Group: group,
            Message: message,
            Members: [user],
            RelationType: 0,
            MessageLimit: 0,
            Action: ''
        },

        User: user,
        ContactList: [user],
        GroupList: [group],
        Message: [message],
        Recipients: [user],
        Status: null,
        Action: '',
        ws: null, // Our websocket
        UsersFromServer: {},
        RecContents: {},
        RecContent: '',
        joined: false, // True if email and username have been filled in
        OnlineUsers: '',
        OnlineUsersList: [],
        OurUsername: ''
    },

    created: function() {
        var element = document.getElementById('chat-messages');
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', e => {
            var msg = JSON.parse(e.data);
            if (typeof msg.User.ID != 'undefined') {
                this.User.ID = msg.User.ID;
            }
            if (msg.Action == 'LoginUser') {
                if (msg.GroupList) {
                    for (var i = 0; i < msg.GroupList.length; i++) {
                        for (
                            var c = 0;
                            c < msg.GroupList[i].Members.length;
                            c++
                        ) {
                            if (
                                msg.GroupList[i].Members[c].Login !=
                                this.OurUsername
                            ) {
                                this.OnlineUsersList.push({
                                    id: msg.GroupList[i].GroupName,
                                    content: msg.GroupList[i].Members[c].Login,
                                    action: this.changeUser
                                });
                            }
                        }
                        if (typeof msg.GroupList[i].Messages != 'undefined') {
                            for (
                                var j = 0;
                                j < msg.GroupList[i].Messages.length;
                                j++
                            ) {
                                if (
                                    typeof this.RecContents[
                                        msg.GroupList[i].GroupName
                                    ] == 'undefined'
                                ) {
                                    this.RecContents[
                                        msg.GroupList[i].GroupName
                                    ] =
                                        '';
                                }
                                this.RecContents[msg.GroupList[i].GroupName] +=
                                    '<div class="chip">' +
                                    msg.GroupList[i].Members[0].Username +
                                    '</div>' +
                                    '<div class="white-text">' +
                                    msg.GroupList[i].Messages[j].Content +
                                    '</div>' +
                                    '<br/>';
                            }
                        }
                    }
                }
                /* TODO: this doesn't work
            } else if (msg.Action == 'SendMessageTo') {
                if (
                    typeof this.RecContents[msg.GroupList[0].GroupName] ==
                    'undefined'
                ) {
                    this.RecContents[msg.GroupList[0].GroupName] = '';

                    for (let i in this.OnlineUsersList) {
                        if (
                            this.OnlineUsersList[i].id ==
                            test.User.Username + msg.Message.User.Username
                        ) {
                            this.OnlineUsers[i].action = changeUser;
                            this.OnlineUsers[i].id = msg.GroupList[0].GroupName;
                            this.OnlineUsers[i].content =
                                msg.Message.User.Username;
                        }
                    }
                }

                this.RecContents[msg.GroupList[0].GroupName] +=
                    '<div class="chip">' +
                    msg.Message.User.Username +
                    '</div>' +
                    '<div class="white-text">' +
                    msg.Message.Content +
                    '</div>' +
                    '<br/>';
                this.RecContent = this.RecContents[msg.GroupList[0].GroupName];
                */
            } else if (msg.Action == 'GetUsers') {
                test.User = msg.User;
                this.OnlineUsers = '';
                this.OnlineUsersList = [];

                /* TODO: 
                * fix Cannot read property 'length' of undefined at WebSocket.ws.addEventListener.e
                */
                // for (var i = 0; i < msg.ContactList.length; i++) {
                //     if (test.User.Username != msg.ContactList[i].Username) {
                //         var gName =
                //             msg.User.Username + msg.ContactList[i].Username;
                //         test.UsersFromServer[gName] = msg.ContactList[i];
                //         this.OnlineUsersList.push({
                //             id: msg.User.Username + msg.ContactList[i].Username,
                //             content: msg.ContactList[i].Username,
                //             action: createGroup
                //         });
                //     }
                // }
            }

            // element.scrollTop = element.scrollHeight; // This doesn't work in this way
        });
    },

    methods: {
        send: function() {
            if (this.MessageIn.Message.Content != '') {
                this.MessageIn.Message.User.Username = this.OurUsername;
                this.MessageIn.Message.MessageSenderID = this.User.ID;
                this.MessageIn.User.Login = this.OurUsername;
                this.MessageIn.User.Username = this.OurUsername;
                this.MessageIn.Message.Content = $('<p>')
                    .html(this.MessageIn.Message.Content)
                    .text();
                this.MessageIn.Message.Group.GroupName = this.MessageIn.Group.GroupName;
                this.MessageIn.Action = 'SendMessageTo';
                if (
                    typeof this.RecContents[this.MessageIn.Group.GroupName] ==
                    'undefined'
                ) {
                    this.RecContents[this.MessageIn.Group.GroupName] = '';
                }
                this.RecContents[this.MessageIn.Group.GroupName] +=
                    '<div class="chip">' +
                    this.OurUsername +
                    '</div>' +
                    '<div class="white-text">' +
                    this.MessageIn.Message.Content +
                    '</div>' +
                    '<br/>';
                var self = this;
                self.RecContent = this.RecContents[
                    this.MessageIn.Group.GroupName
                ];
                this.ws.send(JSON.stringify(this.MessageIn));
                this.MessageIn.Message.Content = '';
            }
        },

        join: function() {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a username', 2000);
                return;
            }
            this.MessageIn.User.Username = $('<p>')
                .html(this.MessageIn.User.Login)
                .text();
            this.MessageIn.User.Login = $('<p>')
                .html(this.MessageIn.User.Login)
                .text();
            this.MessageIn.User.Password = $('<p>')
                .html(this.MessageIn.User.Password)
                .text();
            this.MessageIn.User.Email = 'email';
            this.MessageIn.User.Status = true;
            this.MessageIn.User.UserIcon = 'usericon';
            this.MessageIn.Action = 'LoginUser';
            this.OurUsername = this.MessageIn.User.Username;
            this.ws.send(JSON.stringify(this.MessageIn));
            this.joined = true;
        },

        signUp: function() {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a username', 2000);
                return;
            }
            this.MessageIn.User.Username = $('<p>')
                .html(this.MessageIn.User.Login)
                .text();
            this.MessageIn.User.Login = $('<p>')
                .html(this.MessageIn.User.Login)
                .text();
            this.MessageIn.User.Password = $('<p>')
                .html(this.MessageIn.User.Password)
                .text();
            this.MessageIn.User.Email = 'email';
            this.MessageIn.User.Status = true;
            this.MessageIn.User.UserIcon = 'usericon';
            this.MessageIn.Action = 'CreateUser';
            this.ws.send(JSON.stringify(this.MessageIn));
        },

        showUsers: function() {
            this.MessageIn.User.Login = this.OurUsername;
            this.MessageIn.User.Username = this.OurUsername;
            this.MessageIn.Action = 'GetUsers';
            this.ws.send(JSON.stringify(this.MessageIn));
        },

        changeUser: function(el) {
            this.MessageIn.Group.GroupName = el.id;
            this.RecContent = this.RecContents[el.id];
        },

        createGroup: function(el) {
            this.MessageIn.Action = 'CreateGroup';
            this.MessageIn.Group.GroupTypeID = 1;
            this.MessageIn.Group.User = this.User;
            this.MessageIn.Group.GroupOwnerID = this.User.ID;
            this.MessageIn.Group.GroupName = el.id;
            this.MessageIn.Members[0] = this.User;
            this.MessageIn.Members[1] = this.UsersFromServer[el.id];
            this.ws.send(JSON.stringify(this.MessageIn));

            for (let i in this.OnlineUsersList) {
                if (this.OnlineUsersList[i].id == el.id) {
                    this.OnlineUsers[i].action = this.changeUser;
                    this.OnlineUsers[i].id = this.MessageIn.Group.GroupName;
                    this.OnlineUsers[
                        i
                    ].content = this.MessageIn.Members[1].Username;
                }
            }
        }
    }
});

// function changeUser(el) {
//     test.MessageIn.Group.GroupName = el.id;
//     test.RecContent = test.RecContents[el.id];
// }

// function createGroup(el) {
//     test.MessageIn.Action = 'CreateGroup';
//     test.MessageIn.Group.GroupTypeID = 1;
//     test.MessageIn.Group.User = test.User;
//     test.MessageIn.Group.GroupOwnerID = test.User.ID;
//     test.MessageIn.Group.GroupName = el.id;
//     test.MessageIn.Members[0] = test.User;
//     test.MessageIn.Members[1] = test.UsersFromServer[el.id];
//     test.ws.send(JSON.stringify(test.MessageIn));

//     for (let i in self.OnlineUsersList) {
//         if (self.OnlineUsersList[i].id == el.id) {
//             self.OnlineUsers[i].action = changeUser;
//             self.OnlineUsers[i].id = test.MessageIn.Group.GroupName;
//             self.OnlineUsers[i].content = test.MessageIn.Members[1].Username;
//         }
//     }
// }
