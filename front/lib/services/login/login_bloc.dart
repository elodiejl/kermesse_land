import 'package:bloc/bloc.dart';
import 'package:flutter/foundation.dart';
import '../../models/user_model.dart';
import '../authentication_service.dart';
import 'login_event.dart';
import 'login_state.dart';

class LoginBloc extends Bloc<LoginEvent, LoginState> {
  final AuthenticationService _authService;

  LoginBloc(this._authService) : super(LoginInitial()) {
    on<LoginButtonPressed>(_onLoginButtonPressed);
  }

  void _onLoginButtonPressed(LoginButtonPressed event, Emitter<LoginState> emit) async {
    emit(LoginLoading());
    try {
      if (event.email.isEmpty || event.password.isEmpty) {
        emit(LoginFailure(error: "Email and password cannot be empty"));
        return;
      }

      if (kDebugMode) {
        print('Attempting login with email: ${event.email}');
      }
      bool isLoggedIn = await _authService.login(event.email, event.password);
      if (isLoggedIn) {
        if (kDebugMode) {
          print('Login successful, token: ${_authService.token}');
        }
        emit(LoginSuccess(token: _authService.token));
      } else {
        if (kDebugMode) {
          print('Login failed: Invalid email or password');
        }
        emit(LoginFailure(error: "Invalid email or password"));
      }
    } catch (error) {
      if (kDebugMode) {
        print('Login error: $error');
      }
      emit(LoginFailure(error: "An error occurred: $error"));
    }
  }

  Future<User?> getUserDetails(String token) async {
    try {
      final user = await _authService.getUserDetails(token);
      if (kDebugMode) {
        print('Fetched user details: ${user?.toJson()}');
      }
      return user;
    } catch (error) {
      if (kDebugMode) {
        print('Error fetching user details: $error');
      }
      return null;
    }
  }
}
