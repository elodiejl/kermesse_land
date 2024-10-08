import 'package:equatable/equatable.dart';

abstract class RegistrationEvent extends Equatable {
  @override
  List<Object?> get props => [];
}

class SignUpButtonPressed extends RegistrationEvent {
  final String username;
  final String lastName;
  final String firstName;
  final String email;
  final String password;
  final String roles;

  SignUpButtonPressed({required this.username, required this.lastName, required this.firstName, required this.email, required this.password, required this.roles});

  @override
  List<Object?> get props => [username, lastName, firstName, email, password, roles];

}
