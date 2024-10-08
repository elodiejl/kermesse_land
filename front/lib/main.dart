import 'dart:convert';

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'package:kermesse_land/screens/profile/profile_screen.dart';
import 'package:kermesse_land/screens/login/login_screen.dart';
import 'package:kermesse_land/screens/app.dart';

import 'package:kermesse_land/services/authentication/authentication_bloc.dart';
import 'package:kermesse_land/services/authentication/authentication_state.dart';


import 'package:kermesse_land/utils/config_io.dart';
import 'firebase_options.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:json_theme/json_theme.dart';
import 'package:kermesse_land/utils/routes.dart';
import 'package:provider/provider.dart';
import 'services/authentication_service.dart';
import 'package:rxdart/rxdart.dart';
import 'package:flutter_stripe/flutter_stripe.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  Stripe.publishableKey = 'pk_test_51Q4RUo2M3RaTA7b7Fksh3vTdFqepe6JiALWpSqnYL2oZYuPbwYcwMcUNq0VViATud3Pi61g8BAOcXnBvNK3lym1E00QBxJ82o6';
  try {
    final fileContent = await rootBundle.loadString('.env');
    if (kDebugMode) {
      print('File content: $fileContent');
    }
  } catch (e) {
    if (kDebugMode) {
      print('Error reading file: $e');
    }
  }
  await dotenv.load(fileName: ".env");
  FlutterNativeSplash.preserve(widgetsBinding: WidgetsFlutterBinding.ensureInitialized());

  await Future.delayed(const Duration(seconds: 2));
  FlutterNativeSplash.remove();
  await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);

  final messaging = FirebaseMessaging.instance;

  final settings = await messaging.requestPermission(
    alert: true,
    announcement: false,
    badge: true,
    carPlay: false,
    criticalAlert: false,
    provisional: false,
    sound: true,
  );

  if (kDebugMode) {
    print('Permission granted: ${settings.authorizationStatus}');
  }

  String? token = await messaging.getToken();

  if (kDebugMode) {
    print('Registration Token=$token');
  }

  FirebaseMessaging.onBackgroundMessage(_firebaseMessagingBackgroundHandler);

  final themeStr = await rootBundle.loadString('assets/theme.json');
  final theme = ThemeDecoder.decodeThemeData(jsonDecode(themeStr))!;

  runApp(MyApp(theme: theme));
}

class MyApp extends StatelessWidget {
  final ThemeData theme;

  const MyApp({super.key, required this.theme});

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        Provider<AuthenticationService>(
          create: (_) => AuthenticationService(),
        ),
        BlocProvider(
          create: (context) => AuthenticationBloc(context.read<AuthenticationService>()),
        ),
        ...Config.blocProviders,
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: "Kermesse land Corporation",
        //localizationsDelegates: Config.localizationsDelegates,
        initialRoute: '/',
        routes: getApplicationRoutes(),
        onGenerateRoute: (settings) {
          if (settings.name == '/profile') {
            final String token = settings.arguments as String;
            return MaterialPageRoute(
              builder: (context) => ProfileScreen(token: token),
            );
          }
          return unknownRoute(settings);
        },
        onUnknownRoute: unknownRoute,
        supportedLocales: const [
          Locale('en', ''),
          Locale('fr', ''),
        ],
        theme: theme,
        builder: (context, child) {
          return BlocListener<AuthenticationBloc, AuthenticationState>(
            listener: (context, state) {
              if (state is Unauthenticated) {
                Navigator.of(context).pushReplacementNamed('/login');
              }
            },
            child: child,
          );
        },
        home: _buildHomeScreen(),
      ),
    );
  }

  Widget _buildHomeScreen() {
    return StreamBuilder<User?>(
      stream: FirebaseAuth.instance.authStateChanges(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.active) {
          if (snapshot.hasData) {
            return MainScreen(token: snapshot.data!.uid);
          } else {
            // Redirect to AdminLoginPage if on web, else to LoginPage
            return const LoginPage();
          }
        }
        return const CircularProgressIndicator();
      },
    );
  }
}

Future<void> _firebaseMessagingBackgroundHandler(RemoteMessage message) async {
  await Firebase.initializeApp();

  if (kDebugMode) {
    print("Handling a background message: ${message.messageId}");
    print('Message data: ${message.data}');
    print('Message notification: ${message.notification?.title}');
    print('Message notification: ${message.notification?.body}');
  }
}