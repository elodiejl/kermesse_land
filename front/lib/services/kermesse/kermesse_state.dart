import 'package:equatable/equatable.dart';
import '../../models/kermesse_model.dart';
import '../../models/user_model.dart';

abstract class KermesseState extends Equatable {
  const KermesseState();

  @override
  List<Object> get props => [];
}

class KermesseInitial extends KermesseState {}

class KermesseLoading extends KermesseState {}

class KermesseLoaded extends KermesseState {
  final List<Kermesse> kermesses;

  const KermesseLoaded(this.kermesses);

  @override
  List<Object> get props => [kermesses];
}

class KermesseAdded extends KermesseState {
  final Kermesse kermesse;

  const KermesseAdded(this.kermesse);

  @override
  List<Object> get props => [kermesse];
}

/*class ParticipantLoaded extends KermesseState {
  final List<User> participants;

  const ParticipantLoaded(this.participants);

  @override
  List<Object> get props => [participants];
}

class ParticipantError extends KermesseState {
  final String message;

  const ParticipantError(this.message);

  @override
  List<Object> get props => [message];
}
class KermesseEnrolled extends KermesseState {}*/

class KermesseError extends KermesseState {
  final String message;

  const KermesseError(this.message);

  @override
  List<Object> get props => [message];
}
