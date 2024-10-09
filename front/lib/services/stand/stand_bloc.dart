import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import '../../models/stand_model.dart';
import 'stand_event.dart';
import 'stand_state.dart';
import '../../utils/config_io.dart';

class StandBloc extends Bloc<StandEvent, StandState> {
  StandBloc() : super(StandInitial()) {
    on<FetchStands>((event, emit) async {
      emit(StandLoading());
      try {
        final response = await http.get(
          Uri.parse('${Config.baseUrl}/stands'),
          headers: {'Authorization': 'Bearer ${event.token}'},
        );

        if (response.statusCode == 200) {
          List<Stand> stands = (json.decode(response.body) as List)
              .map((data) => Stand.fromJson(data))
              .toList();
          emit(StandLoaded(stands));
        } else {
          emit(StandError('Failed to load stands'));
        }
      } catch (e) {
        emit(StandError('An error occurred: $e'));
      }
    });

    on<FetchStandsByKermesse>((event, emit) async {
      emit(StandLoading());
      try {
        final response = await http.get(
          Uri.parse('${Config.baseUrl}/stands/kermesse/${event.kermesseId}'),
          headers: {'Authorization': 'Bearer ${event.token}'},
        );

        if (response.statusCode == 200) {
          List<Stand> stands = (json.decode(response.body) as List)
              .map((data) => Stand.fromJson(data))
              .toList();
          emit(StandLoaded(stands));
        } else {
          emit(StandError('Failed to load stands for kermesse'));
        }
      } catch (e) {
        emit(StandError('An error occurred: $e'));
      }
    });
  }
}
