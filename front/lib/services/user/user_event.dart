import 'package:equatable/equatable.dart';

import '../../models/user_model.dart';

abstract class UserEvent extends Equatable {
  const UserEvent();

  @override
  List<Object> get props => [];
}

class FetchUser extends UserEvent {
  final String token;

  const FetchUser(this.token);

  @override
  List<Object> get props => [token];
}

class UpdateUser extends UserEvent {
  final String token;
  final UserUpdate updatedUser;

  const UpdateUser(this.token, this.updatedUser);

  @override
  List<Object> get props => [token, updatedUser];
}