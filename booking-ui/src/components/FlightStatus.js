function getStatusClass(status) {
    switch (status) {
      case 'scheduled': return 'status-scheduled';
      case 'active': return 'status-active';
      case 'arrived': return 'status-arrived';
      default: return '';
    }
  }