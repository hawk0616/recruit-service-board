import React, { useState, useEffect } from 'react';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHeart } from "@fortawesome/free-solid-svg-icons";
import { useLike } from '@/hooks/useLike';
import LoadingSpinner from './LoadingSpinner';
import { useError } from '@/hooks/useError';

type LikeButtonProps = {
  companyId?: number;
  className?: string;
};

const LikeButton: React.FC<LikeButtonProps> = ({ companyId, className }) => {
  const [liked, setLiked] = useState<boolean | undefined>(false);
  const { createLike, deleteLike, checkLike, loading } = useLike();
  const { ErrorHandling } = useError();

  useEffect(() => {
    const fetchLikeStatus = async () => {
      if (companyId) {
        const isLiked = await checkLike(companyId);
        setLiked(isLiked);
      }
    };

    fetchLikeStatus();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [companyId]);

  const toggleLike = async () => {
    if (!companyId || loading) return;

    try {
      if (liked) {
        await deleteLike(companyId);
      } else {
        await createLike(companyId);
      }
      setLiked(!liked);
    } catch (err: any) {
      ErrorHandling(err.message);
    }
  };

  if (liked === null) {
    return <LoadingSpinner />;
  }

  if (!companyId) {
    return (
      <div className="text-gray-400 cursor-not-allowed">
        <FontAwesomeIcon icon={faHeart} size="2x" />
      </div>
    );
  }

  return (
    <button onClick={toggleLike} className={`like-button-class ${className}`}>
      <FontAwesomeIcon icon={faHeart} size="2x" className={liked ? "text-red-500" : "text-gray-400 hover:text-red-500"} />
    </button>
  );
};

export default LikeButton;


