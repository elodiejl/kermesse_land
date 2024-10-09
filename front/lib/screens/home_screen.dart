import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:kermesse_land/screens/stand/stands_by_kermesse_screen.dart';
//import 'package:intl/intl.dart';

import '../../services/kermesse/kermesse_bloc.dart';
import '../../services/kermesse/kermesse_event.dart';
import '../../services/kermesse/kermesse_state.dart';
import '../models/kermesse_model.dart';
import '../services/stand/stand_bloc.dart';
//import '../widgets/geolocation_button.dart';
//import 'hackathon/kermesse_detail_screen.dart';

class HomeScreen extends StatefulWidget {
  final String token;

  const HomeScreen({super.key, required this.token});

  @override
  HomeScreenState createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen> {
  List<Kermesse> _sortedKermesses = [];
  bool _isLoading = false;
  final bool _errorOccurred = false;

  @override
  void initState() {
    super.initState();
    // Fetch Kermesses when the screen is initialized
    context.read<KermesseBloc>().add(FetchKermesses(widget.token));
  }

  void _handleLocationSortedKermesses(List<Kermesse> sortedKermesses) {
    setState(() {
      _sortedKermesses = sortedKermesses;
    });
  }

  void _handleLoadingStateChanged(bool isLoading) {
    setState(() {
      _isLoading = isLoading;
    });
  }

  String formatDate(String dateString) {
    final date = DateTime.parse(dateString);
    return '${date.day}/${date.month}/${date.year}';
  }

  Map<String, dynamic> getStatus(String dateString, String dateFinaleString) {
    final kermesseDate = DateTime.parse(dateString);
    //final kermesseDateFinale = DateTime.parse(dateFinaleString);
    final currentDate = DateTime.now();
    final difference = kermesseDate.difference(currentDate).inDays;
    //final differencedf = kermesseDateFinale.difference(currentDate).inDays;

    if (difference == 0) {
      return {'text': 'En cours', 'color': Colors.orangeAccent};
    } else if (difference < 0) {
      return {'text': 'Terminé', 'color': Colors.red};
    } else if (difference < 3) {
      return {'text': 'Imminent', 'color': Colors.orange};
    } else {
      return {'text': 'À venir', 'color': Colors.green};
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Accueil"),
        backgroundColor: Colors.cyan,
        /*actions: [
          GeoLocationButton(
            token: widget.token,
            onLocationSortedHackathons: (List<dynamic> hackathons) =>
                _handleLocationSortedHackathons(hackathons.cast<Hackathon>()),
            onLoadingStateChanged: _handleLoadingStateChanged,
          ),
        ],*/
      ),
      body: BlocListener<KermesseBloc, KermesseState>(
        listener: (context, state) {
          if (state is KermesseAdded) {
            // Fetch the updated list of Kermesses
            context.read<KermesseBloc>().add(FetchKermesses(widget.token));
          }
        },
        child: _isLoading
            ? const Center(child: CircularProgressIndicator())
            : BlocBuilder<KermesseBloc, KermesseState>(
          builder: (context, state) {
            if (state is KermesseLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is KermesseLoaded) {
              final kermesses = _errorOccurred || _sortedKermesses.isEmpty
                  ? state.kermesses.toList()
                  : _sortedKermesses.toList();

              if (kDebugMode) {
                print(kermesses);
              }

              return RefreshIndicator(
                onRefresh: () async {
                  context.read<KermesseBloc>().add(FetchKermesses(widget.token));
                },
                child: ListView.builder(
                  itemCount: kermesses.length,
                  itemBuilder: (context, index) {
                    final kermesse = kermesses[index];
                    final id = kermesse.id;
                    final date = kermesse.date;
                    final enddate = kermesse.date;
                    if (kDebugMode) {
                      print('date: $date, enddate: $enddate');
                    }
                    final status = getStatus(date, enddate);
                    final formattedDate = formatDate(date);
                    final formattedEndDate = formatDate(enddate);

                    return Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: Card(
                        elevation: 4.0,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(10.0),
                        ),
                        color: Colors.white,
                        child: InkWell(
                          onTap: () {
                            // Action quand on clique sur une kermesse
                            Navigator.of(context).push(
                              MaterialPageRoute(
                                builder: (context) => BlocProvider(
                                  create: (context) => StandBloc(),
                                  child: StandsByKermesseScreen(token: widget.token,
                                    kermesseId: kermesse.id,),
                                ),
                              ),
                            );
                          },
                          child: Padding(
                            padding: const EdgeInsets.all(10.0),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Image.asset(
                                  'assets/logo.png',
                                  width: double.infinity,
                                  height: 100.0,
                                  fit: BoxFit.cover,
                                ),
                                const SizedBox(height: 10),
                                Row(
                                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                                  children: [
                                    Text(
                                      kermesse.name,
                                      style: const TextStyle(
                                        fontSize: 18.0,
                                        fontWeight: FontWeight.bold,
                                      ),
                                    ),
                                    Container(
                                      padding: const EdgeInsets.symmetric(
                                        horizontal: 8.0,
                                        vertical: 4.0,
                                      ),
                                      decoration: BoxDecoration(
                                        color: status['color'],
                                        borderRadius: BorderRadius.circular(12.0),
                                      ),
                                      child: Text(
                                        status['text'],
                                        style: const TextStyle(
                                          color: Colors.white,
                                          fontSize: 12.0,
                                          fontWeight: FontWeight.bold,
                                        ),
                                      ),
                                    ),
                                  ],
                                ),
                                const SizedBox(height: 5),
                                Text(
                                  kermesse.location,
                                  maxLines: 2,
                                  overflow: TextOverflow.ellipsis,
                                  style: TextStyle(
                                    color: Colors.grey[800],
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ),
                      ),
                    );
                  },
                ),
              );
            } else if (state is KermesseError) {
              return Center(child: Text('Error: ${state.message}'));
            } else {
              return Container();
            }
          },
        ),
      ),
    );
  }
}
