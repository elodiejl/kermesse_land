import 'package:equatable/equatable.dart';

//import '../../models/participation_filter_model.dart';

abstract class KermesseEvent extends Equatable {
  const KermesseEvent();

  @override
  List<Object> get props => [];
}

class FetchKermesses extends KermesseEvent {
  final String token;

  const FetchKermesses(this.token);

  @override
  List<Object> get props => [token];
}

class FetchSingleKermesses extends KermesseEvent {
  final String token;
  final String id;

  const FetchSingleKermesses(this.token, this.id);

  @override
  List<Object> get props => [token, id];
}

class AddKermesse extends KermesseEvent {
  final String token;
  final Map<String, dynamic> kermesseData;

  const AddKermesse(this.token, this.kermesseData);

  @override
  List<Object> get props => [token, kermesseData];
}

class FetchKermesseForUser extends KermesseEvent {
  final String token;

  const FetchKermesseForUser(this.token);

  @override
  List<Object> get props => [token];
}

/*class SearchParticipants extends KermesseEvent {
  final int kermesseId;
  final ParticipationFilter filter;
  final String token;

  const SearchParticipants(this.kermesseId, this.filter, this.token);

  @override
  List<Object> get props => [kermesseId, filter, token];
}
class EnrollInKermesse extends KermesseEvent {
  final String token;
  final String kermesseId;

  const EnrollInKermesse(this.token, this.kermesseId);

  @override
  List<Object> get props => [token, kermesseId];
}*/
