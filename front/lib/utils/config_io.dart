import 'dart:io';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:cloud_functions/cloud_functions.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_storage/firebase_storage.dart';


import '../services/authentication_service.dart';
import '../services/login/login_bloc.dart';
import '../services/parent/parent_bloc.dart';
import '../services/register/register_bloc.dart';
import '../services/transaction/transaction_bloc.dart';

class Config {
  static String baseUrl = "https://kermesse-land-5d1bb3e5d576.herokuapp.com/";
  //Platform.isAndroid ? "http://10.0.2.2:8080" : "http://localhost:8080";

  void configureFirebaseEmulators() {
    final host = Platform.isAndroid ? "https://kermesse-land-5d1bb3e5d576.herokuapp.com/" : "localhost:8080";
    FirebaseAuth.instance.useAuthEmulator(host, 9099);
    FirebaseFirestore.instance.useFirestoreEmulator(host, 8082);
    FirebaseStorage.instance.useStorageEmulator(host, 9199);
    FirebaseFunctions.instance.useFunctionsEmulator(host, 5002);
  }

  static List<BlocProvider> get blocProviders => [
    BlocProvider<LoginBloc>(
      create: (context) => LoginBloc(context.read<AuthenticationService>()),
    ),
    BlocProvider<RegistrationBloc>(
      create: (context) => RegistrationBloc(context.read<AuthenticationService>()),
    ),
    BlocProvider<ParentBloc>(
      create: (context) => ParentBloc(),
    ),
    BlocProvider<TransactionBloc>(
      create: (context) => TransactionBloc(),
    ),
    /*
    BlocProvider<StepBloc>(
      create: (context) => StepBloc(StepRepository()),
    ),
    BlocProvider<TeamBloc>(
      create: (context) => TeamBloc(TeamRepository()),
    ),*/
  ];
}
