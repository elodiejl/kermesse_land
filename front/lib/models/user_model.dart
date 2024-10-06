class User {
  final int id;
  final String username;
  final String lastName;
  final String firstName;
  final String email;
  final int roles;

  User({
    required this.id,
    required this.username,
    required this.lastName,
    required this.firstName,
    required this.email,
    required this.roles,
  });

  factory User.fromJson(Map<String, dynamic> json) {

    return User(
      id: json['id'],
      username: json['username'],
      lastName: json['last_name'],
      firstName: json['first_name'],
      email: json['email'],
      roles: json['roles'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'username': username,
      'last_name': lastName,
      'first_name': firstName,
      'email': email,
      'roles': roles,
    };
  }
}

class UserUpdate {
  final int? id; // Now optional
  final String? username; // Now optional
  final String? lastName;
  final String? firstName;
  final String? email;
  //final int? role;

  UserUpdate({
    this.id,
    this.username,
    this.lastName,
    this.firstName,
    this.email,
    //this.role
  });

  Map<String, dynamic> toJson() {
    return {
      'id': id, // Include ID if available
      'username': username,
      'last_name': lastName,
      'first_name': firstName,
      'email': email,
      //'role': role
    };
  }
}
