import 'package:bloc/bloc.dart';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import '../../models/kermesse_model.dart';
//import '../../models/participation_filter_model.dart';
import '../../models/user_model.dart';
import '../../utils/config.dart';
import 'kermesse_event.dart';
import 'kermesse_state.dart';

class KermesseBloc extends Bloc<KermesseEvent, KermesseState> {
  KermesseBloc() : super(KermesseInitial()) {
    on<FetchKermesses>(_onFetchKermesses);
    on<FetchSingleKermesses>(_onFetchSingleKermesses);
    on<AddKermesse>(_onAddKermesse);
    on<FetchKermesseForUser>(_onFetchKermesseForUser);
    //on<SearchParticipants>(_onSearchParticipants);
    //on<EnrollInKermesse>(_onEnrollInKermesse);
  }

  void _onFetchKermesses(FetchKermesses event, Emitter<KermesseState> emit) async {
    emit(KermesseLoading());
    try {
      final kermesses = await _fetchKermesses(event.token);
      emit(KermesseLoaded(kermesses));
    } catch (e) {
      emit(KermesseError(e.toString()));
    }
  }

  Future<List<Kermesse>> _fetchKermesses(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/kermesses'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body);

      if (kDebugMode) {
        print(data);
      }

      return data.map((item) => Kermesse.fromJson(item)).toList();
    } else {
      throw Exception('Failed to load kermesses');
    }
  }

  void _onFetchSingleKermesses(FetchSingleKermesses event, Emitter<KermesseState> emit) async {
    emit(KermesseLoading());
    try {
      if (kDebugMode) {
        print('Fetching single kermesse for ID: ${event.id}');
      }
      final kermesse = await _fetchSingleKermesse(event.token, event.id);
      if (kDebugMode) {
        print('Fetched kermesse: ${kermesse.name}');
      }
      emit(KermesseLoaded([kermesse]));
      if (kDebugMode) {
        print('KermesseLoaded state emitted');
      }
    } catch (e) {
      emit(KermesseError(e.toString()));
    }
  }

  Future<Kermesse> _fetchSingleKermesse(String token, String id) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/kermesses/$id'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (kDebugMode) {
      print('HTTP response status: ${response.statusCode}');
      print('HTTP response body: ${response.body}');
    }

    if (response.statusCode == 200) {
      dynamic data = jsonDecode(response.body)['data'];
      return Kermesse.fromJson(data);
    } else {
      throw Exception('Failed to load kermesse');
    }
  }

  void _onFetchKermesseForUser(FetchKermesseForUser event, Emitter<KermesseState> emit) async {
    emit(KermesseLoading());
    try {
      final kermesses = await _fetchKermesseForUser(event.token);
      emit(KermesseLoaded(kermesses));
    } catch (e) {
      emit(KermesseError(e.toString()));
    }
  }

  Future<List<Kermesse>> _fetchKermesseForUser(String token) async {
    final response = await http.get(
      Uri.parse('${Config.baseUrl}/kermesses/user'),
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body)['data'];
      return data.map((item) => Kermesse.fromJson(item)).toList();
    } else {
      throw Exception('Failed to load kermesses');
    }
  }

  void _onAddKermesse(AddKermesse event, Emitter<KermesseState> emit) async {
    try {
      final kermesse = await _addKermesse(event.token, event.kermesseData);
      emit(KermesseAdded(kermesse));
      add(FetchKermesseForUser(event.token)); // Fetch the updated list of Kermesses for the user
    } catch (e) {
      if (kDebugMode) {
        print('Error adding kermesse: $e');
      }
      emit(KermesseError(e.toString()));
    }
  }

  Future<Kermesse> _addKermesse(String token, Map<String, dynamic> kermesseData) async {
    final url = Uri.parse('${Config.baseUrl}/kermesses');

    final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode(kermesseData),
    );

    if (kDebugMode) {
      print('Response status: ${response.statusCode}');
      print('Response headers: ${response.headers}');
      print('Response body: ${response.body}');
    }

    if (response.statusCode == 201) {
      return Kermesse.fromJson(jsonDecode(response.body)['data']);
    } else if (response.statusCode == 400) {
      throw Exception('Bad request');
    } else {
      throw Exception('Failed to add kermesse');
    }
  }

  /*void _onSearchParticipants(SearchParticipants event, Emitter<KermesseState> emit) async {
    emit(KermesseLoading());
    try {
      final participants = await _searchParticipants(event.kermesseId, event.filter, event.token);
      emit(ParticipantLoaded(participants));
    } catch (e) {
      if (kDebugMode) {
        print('Error in _onSearchParticipants: $e');
      }
      emit(ParticipantError(e.toString()));
    }
  }

  Future<List<User>> _searchParticipants(int kermesseId, ParticipationFilter filter, String token) async {
    final url = '${Config.baseUrl}/kermesses/$kermesseId/teammate/search';
    final headers = {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $token',
    };
    final body = json.encode(filter.toJson());

    if (kDebugMode) {
      print('Sending request to $url');
      print('Headers: $headers');
      print('Body: $body');
    }

    final response = await http.post(
      Uri.parse(url),
      headers: headers,
      body: body,
      );

    if (kDebugMode) {
      print('Response status: ${response.statusCode}');
        print('Response body: ${response.body}');
    }

    if (response.statusCode == 200) {
      List<dynamic> data = jsonDecode(response.body)['data'];
      return data.map((item) => User.fromJson(item)).toList();
    } else {
      throw Exception('Failed to search participants');
    }
  }
      
  void _onEnrollInKermesse(EnrollInKermesse event, Emitter<KermesseState> emit) async {
    emit(KermesseLoading());
    try {
      final response = await _enrollInKermesse(event.token, event.kermesseId);
      if (response.statusCode == 200) {
        emit(KermesseEnrolled());
      } else {
        emit(const KermesseError('Échec de l\'inscription'));
      }
    } catch (e) {
      emit(KermesseError('Échec de l\'inscription: $e'));
    }
  }

  Future<http.Response> _enrollInKermesse(String token, String KermesseId) async {
    final url = Uri.parse('${Config.baseUrl}/kermesses/$kermesseId/enroll');

    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer $token',
      },
    );

    if (kDebugMode) {
      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');
    }

    return response;
  }*/
}
