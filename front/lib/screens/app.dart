import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:kermesse_land/screens/profile/profile_screen.dart';
import 'package:kermesse_land/services/kermesse/kermesse_event.dart';
import 'home_screen.dart';
import 'package:kermesse_land/services/kermesse/kermesse_bloc.dart';
import '../services/kermesse/kermesse_event.dart';
import 'kermesse/kermesse_screen.dart';

class MainScreen extends StatefulWidget {
  final String token;

  const MainScreen({super.key, required this.token});

  @override
  MainScreenState createState() => MainScreenState();
}

class MainScreenState extends State<MainScreen> {
  int _selectedIndex = 0;
  late List<Widget> _widgetOptions;
  late KermesseBloc _kermesseBloc;

  @override
  void initState() {
    super.initState();
    _kermesseBloc = KermesseBloc();
    _widgetOptions = [
      HomeScreen(token: widget.token),
      ProfileScreen(token: widget.token),
    ];
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });

    // Refresh kermesses data when the user navigates to the home screen
    if (index == 0) {
      _kermesseBloc.add(FetchKermesses(widget.token));
    } else if (index == 1) {
      _kermesseBloc.add(FetchKermesseForUser(widget.token));
    }
  }

  @override
  Widget build(BuildContext context) {
    return BlocProvider.value(
      value: _kermesseBloc,
      child: Scaffold(
        body: IndexedStack(
          index: _selectedIndex,
          children: _widgetOptions.map((widget) => BlocProvider.value(value: _kermesseBloc, child: widget)).toList(),
        ),
        bottomNavigationBar: BottomNavigationBar(
          items: const <BottomNavigationBarItem>[
            BottomNavigationBarItem(
              icon: Icon(Icons.home),
              label: 'Accueil',
            ),
            /*BottomNavigationBarItem(
              icon: Icon(Icons.event),
              label: 'Mes kermesses',
            ),*/
            BottomNavigationBarItem(
              icon: Icon(Icons.person),
              label: 'Profil',
            ),
          ],
          currentIndex: _selectedIndex,
          selectedItemColor: Colors.cyan,
          onTap: _onItemTapped,
        ),
      ),
    );
  }
}
