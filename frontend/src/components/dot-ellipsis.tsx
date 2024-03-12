import React from 'react';

const TripleDotLoader: React.FC = () => {
  return (
    <div className="flex justify-center items-center space-x-1">
      <div className="bounce h-2 w-2 bg-accent-foreground/50 rounded-full"></div>
      <div className="bounce h-2 w-2 bg-accent-foreground/50 rounded-full" style={{ animationDelay: '0.1s' }}></div>
      <div className="bounce h-2 w-2 bg-accent-foreground/50 rounded-full" style={{ animationDelay: '0.2s' }}></div>
    </div>
  );
};

export default TripleDotLoader;
