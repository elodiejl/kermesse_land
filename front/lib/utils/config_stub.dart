// config_stub.dart
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
//import 'package:flutter_localizations/flutter_localizations.dart';
//import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import '../services/authentication/authentication_bloc.dart';
import '../services/authentication_service.dart';
import '../services/login/login_bloc.dart';
import '../services/register/register_bloc.dart';

class Config {
  static String baseUrl = "https://kermesse-land-5d1bb3e5d576.herokuapp.com";

  /*static List<LocalizationsDelegate> get localizationsDelegates => [
    AppLocalizations.delegate,
    GlobalMaterialLocalizations.delegate,
    GlobalWidgetsLocalizations.delegate,
    GlobalCupertinoLocalizations.delegate,
  ];*/

  void configureFirebaseEmulators() {
    // Firebase emulators are not typically used in web environments.
  }

  static List<BlocProvider> get blocProviders => [
    BlocProvider<AuthenticationBloc>(
      create: (context) => AuthenticationBloc(AuthenticationService()),
    ),
    BlocProvider<LoginBloc>(
      create: (context) => LoginBloc(AuthenticationService()),
    ),
    BlocProvider<RegistrationBloc>(
      create: (context) => RegistrationBloc(AuthenticationService()),
    ),
    /*BlocProvider<KermesseBloc>(
      create: (context) => KermesseBloc(),
    ),*/
  ];
}
