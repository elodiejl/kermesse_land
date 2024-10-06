import 'package:flutter_bloc/flutter_bloc.dart';
import 'logout_event.dart';
import 'logout_state.dart';
import '../authentication_service.dart';

class LogoutBloc extends Bloc<LogoutEvent, LogoutState> {
  final AuthenticationService _authenticationService;

  LogoutBloc(this._authenticationService) : super(LogoutInitial()) {
    on<LogoutRequested>((event, emit) async {
      emit(LogoutInProgress());
      try {
        await _authenticationService.logout();
        emit(LogoutSuccess());
      } catch (error) {
        emit(LogoutFailure(error: error.toString()));
      }
    });
  }
}
