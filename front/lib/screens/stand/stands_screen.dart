import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../services/stand/stand_bloc.dart';
import '../../services/stand/stand_event.dart';
import '../../services/stand/stand_state.dart';

class StandsScreen extends StatelessWidget {
  final String token;

  StandsScreen({super.key, required this.token});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Stands')),
      body: BlocProvider(
        create: (context) => StandBloc()..add(FetchStands(token)),
        child: StandList(),
      ),
    );
  }
}

class StandList extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return BlocBuilder<StandBloc, StandState>(
      builder: (context, state) {
        if (state is StandLoading) {
          return const Center(child: CircularProgressIndicator());
        } else if (state is StandLoaded) {
          return ListView.builder(
            itemCount: state.stands.length,
            itemBuilder: (context, index) {
              final stand = state.stands[index];
              return ListTile(
                title: Text(stand.name),
                subtitle: Text('Prix : ${stand.participationCost}, Type: ${stand.standType}'),
              );
            },
          );
        } else if (state is StandError) {
          return Center(child: Text(state.message));
        }
        return const Center(child: Text('Aucun stand trouvé.'));
      },
    );
  }
}
