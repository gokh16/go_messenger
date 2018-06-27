

var user={
    'Login': '',
    'Password': '',
    'Username': '',
    'Email': '',
    'Status': false,
    'UserIcon': '',
};

var message={
    'Content': '',
};
var group={
    'GroupName': '',
    'Messages':[message],
    'Members': [user],
};
new Vue({
    el: '#app',
    data: {
        MessageIn:{
            User: user,
            Contact: user,
            Group: group,
            Message: message,
            Members: [user],
            RelationType: 0,
            MessageLimit: 0,
            Action: '',
        },

        User: user,
        ContactList: [user],
        GroupList: [group],
        Status: null,
        Action: '',
        ws: null, // Our websocket

        RecContents: {},
        RecContent: '',
        joined: false, // True if email and username have been filled in
        OnlineUsers: '',
    },


    created: function () {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);
            if (msg.Action == "LoginUser") {
                for (var i = 0; i < msg.GroupList.length; i++) {
                    for (var c = 0; c < msg.GroupList[i].Members.length; c++) {
                        self.OnlineUsers += '<div class="white-text">' + msg.GroupList[i].Members[c].Username + '</div>' + '<br/>';
                    }
                }
            } else if (msg.Action == "SendMessageTo") {
                self.RecContent +=
                    '<div class="chip">' + msg.User.Username + ' from: ' + msg.GroupList[0].GroupName + '</div>' +
                    '<div class="white-text">' + msg.GroupList[0].Messages.Content + '</div>' +
                    '<br/>';
            }
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight;// Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.MessageIn.Message.Content != '') {
                this.MessageIn.Message.Content = $('<p>').html(this.MessageIn.Message.Content).text();
                this.MessageIn.Action = "SendMessageTo";
                this.ws.send(JSON.stringify(this.MessageIn));
                this.MessageIn.Message.Content = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.MessageIn.User.Login) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.MessageIn.User.Username = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Login = $('<p>').html(this.MessageIn.User.Login).text();
            this.MessageIn.User.Password = $('<p>').html(this.MessageIn.User.Password).text();
            this.MessageIn.User.Email = "email";
            this.MessageIn.User.Status = true;
            this.MessageIn.User.UserIcon = "usericon";
            this.MessageIn.Action = "LoginUser";
            this.ws.send(JSON.stringify(this.MessageIn));
            this.joined = true;
        },
    }
});
