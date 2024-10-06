import 'package:bloc/bloc.dart';
import '../authentication_service.dart';
import 'package:front/services/authentication/authentication_state.dart';

import 'package:front/services/authentication/authentication_event.dart';


class AuthenticationBloc extends Bloc<AuthenticationEvent, AuthenticationState> {
  final AuthenticationService _authenticationService;

  AuthenticationBloc(this._authenticationService) : super(AuthenticationInitial()) {
    on<LogoutEvent>((event, emit) async {
      await _authenticationService.logout();
      emit(Unauthenticated());
    });
  }
}
